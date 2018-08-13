package inode

import (
	"io/ioutil"
	"os"
	"testing"
	"time"

	"golang.org/x/sys/unix"

	"github.com/swiftstack/ProxyFS/conf"
	"github.com/swiftstack/ProxyFS/dlm"
	"github.com/swiftstack/ProxyFS/evtlog"
	"github.com/swiftstack/ProxyFS/headhunter"
	"github.com/swiftstack/ProxyFS/logger"
	"github.com/swiftstack/ProxyFS/ramswift"
	"github.com/swiftstack/ProxyFS/stats"
	"github.com/swiftstack/ProxyFS/swiftclient"
)

func testSetup(t *testing.T, starvationMode bool) {
	var (
		doneChan              chan bool
		err                   error
		signalHandlerIsArmed  bool
		testDir               string
		testConfMap           conf.ConfMap
		testConfStrings       []string
		testConfUpdateStrings []string
	)

	testDir, err = ioutil.TempDir(os.TempDir(), "ProxyFS_test_inode_")
	if nil != err {
		t.Fatalf("ioutil.TempDir() failed: %v", err)
	}

	err = os.Chdir(testDir)
	if nil != err {
		t.Fatalf("os.Chdir() failed: %v", err)
	}

	testConfStrings = []string{
		"Stats.IPAddr=localhost",
		"Stats.UDPPort=52184",
		"Stats.BufferLength=100",
		"Stats.MaxLatency=1s",
		"Logging.LogFilePath=/dev/null",
		"Logging.LogToConsole=false",
		"SwiftClient.NoAuthIPAddr=127.0.0.1",
		"SwiftClient.NoAuthTCPPort=45262",
		"SwiftClient.Timeout=10s",
		"SwiftClient.RetryLimit=3",
		"SwiftClient.RetryLimitObject=3",
		"SwiftClient.RetryDelay=10ms",
		"SwiftClient.RetryDelayObject=10ms",
		"SwiftClient.RetryExpBackoff=1.2",
		"SwiftClient.RetryExpBackoffObject=1.0",
		"FlowControl:TestFlowControl.MaxFlushSize=10000000",
		"FlowControl:TestFlowControl.MaxFlushTime=10s",
		"FlowControl:TestFlowControl.ReadCacheLineSize=1000000",
		"FlowControl:TestFlowControl.ReadCacheWeight=100",
		"PhysicalContainerLayout:PhysicalContainerLayoutReplicated3Way.ContainerStoragePolicy=silver",
		"PhysicalContainerLayout:PhysicalContainerLayoutReplicated3Way.ContainerNamePrefix=Replicated3Way_",
		"PhysicalContainerLayout:PhysicalContainerLayoutReplicated3Way.ContainersPerPeer=1000",
		"PhysicalContainerLayout:PhysicalContainerLayoutReplicated3Way.MaxObjectsPerContainer=1000000",
		"Peer:Peer0.PrivateIPAddr=localhost",
		"Peer:Peer0.ReadCacheQuotaFraction=0.20",
		"Cluster.Peers=Peer0",
		"Cluster.WhoAmI=Peer0",
		"Volume:TestVolume.FSID=1",
		"Volume:TestVolume.PrimaryPeer=Peer0",
		"Volume:TestVolume.AccountName=AUTH_test",
		"Volume:TestVolume.CheckpointContainerName=.__checkpoint__",
		"Volume:TestVolume.CheckpointContainerStoragePolicy=gold",
		"Volume:TestVolume.CheckpointInterval=10s",
		"Volume:TestVolume.DefaultPhysicalContainerLayout=PhysicalContainerLayoutReplicated3Way",
		"Volume:TestVolume.FlowControl=TestFlowControl",
		"Volume:TestVolume.NonceValuesToReserve=100",
		"Volume:TestVolume.MaxEntriesPerDirNode=32",
		"Volume:TestVolume.MaxExtentsPerFileNode=32",
		"Volume:TestVolume.MaxInodesPerMetadataNode=32",
		"Volume:TestVolume.MaxLogSegmentsPerMetadataNode=64",
		"Volume:TestVolume.MaxDirFileNodesPerMetadataNode=16",
		"Volume:TestVolume.MaxBytesInodeCache=100000",
		"Volume:TestVolume.InodeCacheEvictInterval=1s",
		"FSGlobals.VolumeList=TestVolume",
		"FSGlobals.InodeRecCacheEvictLowLimit=10000",
		"FSGlobals.InodeRecCacheEvictHighLimit=10010",
		"FSGlobals.LogSegmentRecCacheEvictLowLimit=10000",
		"FSGlobals.LogSegmentRecCacheEvictHighLimit=10010",
		"FSGlobals.BPlusTreeObjectCacheEvictLowLimit=10000",
		"FSGlobals.BPlusTreeObjectCacheEvictHighLimit=10010",
		"FSGlobals.DirEntryCacheEvictLowLimit=10000",
		"FSGlobals.DirEntryCacheEvictHighLimit=10010",
		"FSGlobals.FileExtentMapEvictLowLimit=10000",
		"FSGlobals.FileExtentMapEvictHighLimit=10010",
		"RamSwiftInfo.MaxAccountNameLength=256",
		"RamSwiftInfo.MaxContainerNameLength=256",
		"RamSwiftInfo.MaxObjectNameLength=1024",
		"RamSwiftInfo.AccountListingLimit=10000",
		"RamSwiftInfo.ContainerListingLimit=10000",
	}

	testConfMap, err = conf.MakeConfMapFromStrings(testConfStrings)
	if nil != err {
		t.Fatalf("conf.MakeConfMapFromStrings() failed: %v", err)
	}

	if starvationMode {
		testConfUpdateStrings = []string{
			"SwiftClient.ChunkedConnectionPoolSize=1",
			"SwiftClient.NonChunkedConnectionPoolSize=1",
		}
	} else {
		testConfUpdateStrings = []string{
			"SwiftClient.ChunkedConnectionPoolSize=512",
			"SwiftClient.NonChunkedConnectionPoolSize=128",
		}
	}

	err = testConfMap.UpdateFromStrings(testConfUpdateStrings)
	if nil != err {
		t.Fatalf("testConfMap.UpdateFromStrings(testConfUpdateStrings) failed: %v", err)
	}

	signalHandlerIsArmed = false
	doneChan = make(chan bool, 1)
	go ramswift.Daemon("/dev/null", testConfStrings, &signalHandlerIsArmed, doneChan, unix.SIGTERM)

	for !signalHandlerIsArmed {
		time.Sleep(100 * time.Millisecond)
	}

	err = logger.Up(testConfMap)
	if nil != err {
		t.Fatalf("logger.Up() failed: %v", err)
	}

	err = evtlog.Up(testConfMap)
	if nil != err {
		t.Fatalf("evtlog.Up() failed: %v", err)
	}

	err = stats.Up(testConfMap)
	if nil != err {
		t.Fatalf("stats.Up() failed: %v", err)
	}

	err = dlm.Up(testConfMap)
	if nil != err {
		t.Fatalf("dlm.Up() failed: %v", err)
	}

	err = swiftclient.Up(testConfMap)
	if nil != err {
		t.Fatalf("swiftclient.Up() failed: %v", err)
	}

	err = headhunter.Format(testConfMap, "TestVolume")
	if nil != err {
		t.Fatalf("headhunter.Format() failed: %v", err)
	}

	err = headhunter.Up(testConfMap)
	if nil != err {
		t.Fatalf("headhunter.Up() failed: %v", err)
	}

	err = Up(testConfMap)
	if nil != err {
		t.Fatalf("inode.Up() failed: %v", err)
	}
}

func testTeardown(t *testing.T) {
	var (
		err     error
		testDir string
	)

	err = Down()
	if nil != err {
		t.Fatalf("inode.Down() failed: %v", err)
	}

	err = headhunter.Down()
	if nil != err {
		t.Fatalf("headhunter.Down() failed: %v", err)
	}

	err = swiftclient.Down()
	if nil != err {
		t.Fatalf("swiftclient.Down() failed: %v", err)
	}

	err = dlm.Down()
	if nil != err {
		t.Fatalf("dlm.Down() failed: %v", err)
	}

	err = stats.Down()
	if nil != err {
		t.Fatalf("stats.Down() failed: %v", err)
	}

	err = evtlog.Down()
	if nil != err {
		t.Fatalf("evtlog.Down() failed: %v", err)
	}

	err = logger.Down()
	if nil != err {
		t.Fatalf("logger.Down() failed: %v", err)
	}

	testDir, err = os.Getwd()
	if nil != err {
		t.Fatalf("os.Getwd() failed: %v", err)
	}

	err = os.Chdir("..")
	if nil != err {
		t.Fatalf("os.Chdir() failed: %v", err)
	}

	err = os.RemoveAll(testDir)
	if nil != err {
		t.Fatalf("os.RemoveAll() failed: %v", err)
	}
}
