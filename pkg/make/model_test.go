package make_test

import (
	"dtweave.io/zookeeper-operator/api/v1alpha1"
	makeutil "dtweave.io/zookeeper-operator/pkg/make"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	v1 "k8s.io/api/core/v1"
	v12 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var _ = Describe("Configmap make model", func() {
	var cm *v1.ConfigMap
	var err error
	var zookeeper v1alpha1.Zookeeper
	var cfg string

	BeforeEach(func() {
		zookeeper = v1alpha1.Zookeeper{
			ObjectMeta: v12.ObjectMeta{
				Name:      "zookeeper-cfg-cm",
				Namespace: "default",
			},
			Spec: v1alpha1.ZookeeperSpec{
				Conf: v1alpha1.ZookeeperConf{
					JvmFlags:               "",
					LogLevel:               v1alpha1.ERROR,
					DataDir:                "/data",
					DataLogDir:             "/log",
					TickTime:               2000,
					InitLimit:              10,
					SyncLimit:              2,
					PreAllocSize:           65536,
					SnapCount:              100000,
					MaxClientCnxns:         60,
					MaxSessionTimeout:      40000,
					MinSessionTimeout:      4000,
					GlobalOutstandingLimit: 1000,
					CommitLogCount:         500,
					SnapSizeLimitInKb:      4194304,
					QuorumListenOnAllIPs:   false,
					ExistingCfgConfigmap:   "",
					Autopurge: v1alpha1.Autopurge{
						SnapRetainCount: 3,
						PurgeInterval:   0,
					},
				},
			},
		}
		cm, err = makeutil.Configmap(&zookeeper)
		cfg = cm.Data["zoo.cfg"]
	})
	Context("configmap", func() {
		It("err is nil", func() {
			Ω(err).To(BeNil())
		})
	})
	Context("zoo.cfg", func() {
		It("should have a data dir", func() {
			Ω(cfg).To(ContainSubstring("dataDir=/data"))
		})
		It("should have a data log dir", func() {
			Ω(cfg).To(ContainSubstring("dataLogDir=/log"))
		})
		It("should set tickTime is 2000", func() {
			Ω(cfg).To(ContainSubstring("tickTime=2000"))
		})
		It("should set autopurge.snapRetainCount is 3", func() {
			Ω(cfg).To(ContainSubstring("autopurge.snapRetainCount=3"))
		})
		It("should set autopurge.purgeInterval is0", func() {
			Ω(cfg).To(ContainSubstring("autopurge.purgeInterval=0"))
		})
	})
})
