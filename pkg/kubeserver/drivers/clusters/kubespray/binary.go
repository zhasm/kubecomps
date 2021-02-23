package kubespray

import "yunion.io/x/pkg/errors"

const (
	kubeadm = "kubeadm"
	kubelet = "kubelet"
	kubectl = "kubectl"
	kubecni = "kubecni"
	etcd    = "etcd"
	helm    = "helm"
	amd64   = "amd64"
	arm64   = "arm64"
)

const (
	AMD64 = amd64
	ARM64 = arm64
)

// code copy from: https://github.com/kubesphere/kubekey/blob/d2a78d20c4a47ab55501ac65f11d54ae51514b1f/pkg/files/file.go#L40
var (
	fileSha256 = map[string]map[string]map[string]string{
		kubeadm: {
			amd64: {
				"v1.15.12": "e052bae41e731921a9197b4d078c30d33fac5861716dc275bfee4670addbac9b",
				"v1.16.8":  "58a74986af13b969abc8b471822f36f3fda71f95ed1c006f48c8d2ab88f8edf1",
				"v1.16.10": "726d42c569f25078d03b758477f17f543c845aef2ff48acd9d4269705ca1aa9d",
				"v1.16.12": "bb4d0f045600b883745016416c14533f823d582f4f20df691b7f79a6545b6480",
				"v1.16.13": "3ddce3fb919f1e8b0a3e0a1ae1d20c9af0fd4a7d731be1e818597b3ecdb49023",
				"v1.17.0":  "0d8443f50fb7caab2e5e7e53f9dc56d5ffe55f021ec061f2e2bcba0481df5a48",
				"v1.17.4":  "3cdcffcf8a1660241a045cfdfed3ebbf7f7c6a0840f008e2b049b533bca5bb8c",
				"v1.17.5":  "9bd2fd1118b3d07d12e2a806c04bf34d99e79886c5318ddc003ba38f30da390c",
				"v1.17.6":  "d4cfc9a0a734ba015594974ee4253b8965b95cdb6e83d8a6a946675aad418b40",
				"v1.17.7":  "9d4b97e93ddb204798b91fec063743e218c92b42798779b5248a49e1476226e2",
				"v1.17.8":  "c59b85696c4cbabe896ba71f4bbc99e4ad2444fcea851e3ee740705584420aad",
				"v1.17.9":  "5ef1660d3d56e93e3d87d6a7028aa64745984be0b0678c45c32f66043b4d69b4",
				"v1.18.3":  "a60974e9840e006076d204fd4ddcba96213beba10fb89ff01882095546c9684d",
				"v1.18.5":  "e428fc9d1cf860090346a83eb66082c3be6b6032f0db9e4f8e6d52492d46231f",
				"v1.18.6":  "11b4180b9f82a8b6bb30250e3d7341b104521f3b654076b8569853ec9451b2a9",
				"v1.18.8":  "27c8f4d4398d57762998b157d35802a36a7ea9b2b6f9a363c397a9d65b4f3c89",
				"v1.19.0":  "88ce7dc5302d8847f6e679aab9e4fa642a819e8a33d70731fb7bc8e110d8659f",
			},
			arm64: {
				"v1.15.12": "dfc1af35cccac89099a7e9a48dcc4b0d956a8b1d4dfcdd2be12191b6f6c384a3",
				"v1.16.8":  "2300e2a7dc16512595c7aebc486799239039d33f33db2d085550d1f2d5f3129b",
				"v1.16.12": "67f675f8fb1ff3af56ca0a976323a65cabc35efa53b7896146684b8f53990741",
				"v1.16.13": "bb4d0f045600b883745016416c14533f823d582f4f20df691b7f79a6545b6480",
				"v1.17.0":  "0b94d1ace240a8f9995358ca2b66ac92072e3f3cd0543275b315dcd317798546",
				"v1.17.7":  "6c8622adf5a7a2dfc66ebe15058353b2e2660b01f1e8990bab7a9c7fca76bccb",
				"v1.17.8":  "5a52e7d0306890e68ed66fc47ecd70bf14628c70527442fd0cd2973dbde7064c",
				"v1.17.9":  "b56dc03177636fdafb4f8ab329d087b804cb7395c142f76e8246e86083c6d750",
				"v1.18.5":  "0e2a9de622177015c2514498382b0d821ac8f71c7ed5f02e5684d456ff3c0e4d",
				"v1.18.6":  "df5a3d7c70c3f8221d57093c5cb17558aad6e65725d7a096c6620302fbf64730",
				"v1.18.8":  "71f6d95f165a9e8066c6f299217af779829ab3d798f6130caf6daa4784dc0464",
				"v1.19.0":  "db1c432646e6e6484989b6f7191f3610996ac593409f12574290bfc008ea11f5",
			},
		},
		kubelet: {
			amd64: {
				"v1.15.12": "dff48393a3116b8f7dea206b81678e52f7fad298f1aff976f18f1bfa4e9ccdde",
				"v1.16.8":  "4573da19fed14c84f4434ab7cbedf5ded4bf89710c078d58c0703cf2332df198",
				"v1.16.10": "82b38f444d11c2436040165b1addf46d0909a6daec9133cc979678835ef8e14b",
				"v1.16.12": "fbc8c16b148dbb3234a3e13f80e6c6736557c10f8c046edfb1dc5337fe2dd40f",
				"v1.16.13": "a88c0e9f8c4b5a2e91c2c4a8d772cc65ca3a0eb5d477cbce06fbf82d3e50c158",
				"v1.17.0":  "c2af77f501c3164e80171903028d35c632366f53dec0c8419828d4e55d86146f",
				"v1.17.4":  "f3a427ddf610b568db60c8d47565041901220e1bbe257614b61bb4c76801d765",
				"v1.17.5":  "c5fbfa83444bdeefb51934c29f0b4b7ffc43ce5a98d7f957d8a11e3440055383",
				"v1.17.6":  "4b7fd5123bfafe2249bf91ed83469c2655a8d3295966e5fbd952f89b64b75f57",
				"v1.17.7":  "a6b66c94a37dd6ae830a9af5b9200884a2c0af868096a3c2553b2e876723c2a2",
				"v1.17.8":  "b39081fb40332ae12d262b04dc81630e5c6550fb196f09b60f3d726283dff17f",
				"v1.17.9":  "3b6cdfcd38a646c7b553821ef9bb67e93541da658305c00705e6ab2ba15e73af",
				"v1.18.3":  "6aac8853028a4f185de5ccb5b41b3fbd87726161445dee56f351e3e51442d669",
				"v1.18.5":  "8c328f65d30f0edd0fd4f529b09d6fc588cfb7b524d5c9f181e36de6e494e19c",
				"v1.18.6":  "2eb9baf5a65a7b94c653dbd7af03a768a520961eb27ef369e43ef12711e22d4a",
				"v1.18.8":  "a4116675ac52bf80e224fba8ff6db6f2d7aed192bf6fffd5f8e4d5efb4368f31",
				"v1.19.0":  "3f03e5c160a8b658d30b34824a1c00abadbac96e62c4d01bf5c9271a2debc3ab",
			},
			arm64: {
				"v1.15.12": "c7f586a77acdb3c3e27a6b3bd749760538b830414575f8718f03f7ce53b138d8",
				"v1.16.8":  "a6889c9957d8ec3ba15676b1e2eff021c9d120284f185d367626763dd15a245b",
				"v1.16.12": "0ef9d42e27bf85e9ff276f2181e17e2912941c3a7ae9086de722ac3c9cea997f",
				"v1.16.13": "bb4d0f045600b883745016416c14533f823d582f4f20df691b7f79a6545b6480",
				"v1.17.0":  "b1a4a2325383854a69ec768e7dc00f69378d3ccbc554859d910bf5b582264ea2",
				"v1.17.7":  "eb1715a745281f6aee34644653f73787acdd9f3904e3d58e1319ded4a16be013",
				"v1.17.8":  "673355f62aa422915682ae595e4e53813e4656f2c272eb032f97492211cfced5",
				"v1.17.9":  "d57c25a3d67c937a9d6778de07295478185f73938937868525030a01d15c372f",
				"v1.18.5":  "c3815bc740755aa9fd3ec240ad808a13628a4deb6ec2b4338e772fd0cf77e1a2",
				"v1.18.6":  "257fd42be375025fb93724bda9bef23b73eb40531f22bab9e19f6d6ff1ca57cf",
				"v1.18.8":  "d36e2d656bad232e8b48b19c948164ee3966669f4566cf5ea43ca22f6eed1aa5",
				"v1.19.0":  "d8fa5a9739ecc387dfcc55afa91ac6f4b0ccd01f1423c423dbd312d787bbb6bf",
			},
		},
		kubectl: {
			amd64: {
				"v1.15.12": "a32b762279c33cb8d8f4198f3facdae402248c3164e9b9b664c3afbd5a27472e",
				"v1.16.8":  "1d8602496ca4b843824a9746206509991eb8d30b5bb8436b36a02718729934ed",
				"v1.16.10": "246d36e4ce67e74e95ff2ba578b9189f58e5def0e8830a24cd30fa3cf279742f",
				"v1.16.12": "db72e5c90de59e1bf287bef55eaf0b603c8d74b3dc552f356ccc02b08c2eb348",
				"v1.16.13": "ab861ec3ec347062bd1b87f8d78d15cd1ce251e74c5fe662e434056962d2a2c9",
				"v1.17.0":  "6e0aaaffe5507a44ec6b1b8a0fb585285813b78cc045f8804e70a6aac9d1cb4c",
				"v1.17.4":  "465b2d2bd7512b173860c6907d8127ee76a19a385aa7865608e57a5eebe23597",
				"v1.17.5":  "03cd1fa19f90d38005148793efdb17a9b58d01dedea641a8496b9cf228db3ab4",
				"v1.17.6":  "5e245f6af6fb761fbe4b3ac06b753f33b361ce0486c48c85b45731a7ee5e4cca",
				"v1.17.7":  "7124a296518edda2ae326e754aec9be6d0ac86131e6f61b52f5ecaa413b66ae4",
				"v1.17.8":  "01283cbc2b09555cbf2a71c162097552a62a4fd48a0a4c06e34e9b853b815486",
				"v1.17.9":  "2ca83eecd221bedf3eceb0ccfcf45bb2e27950c382c2326211303adb0a9c4232",
				"v1.18.3":  "6fcf70aae5bc64870c358fac153cdfdc93f55d8bae010741ecce06bb14c083ea",
				"v1.18.5":  "69d9b044ffaf544a4d1d4b40272f05d56aaf75d7e3c526d5418d1d3c78249e45",
				"v1.18.6":  "62fcb9922164725c7cba5747562f2ad2f4d834ad0a458c1e4c794cc203dcdfb3",
				"v1.18.8":  "a076f5eff0710de94d1eb77bee458ea43b8f4d9572bbb3a3aec1edf0dde0a3e7",
				"v1.19.0":  "79bb0d2f05487ff533999a639c075043c70a0a1ba25c1629eb1eef6ebe3ba70f",
			},
			arm64: {
				"v1.15.12": "ef9a4272d556851c645d6788631a2993823260a7e1176a281620284b4c3406da",
				"v1.16.8":  "d08aab5f02db63690672e5d9052659589301323c010d90734788d5332ac99daa",
				"v1.16.12": "7f493dcf9d4edfeea68284c4cd7c74383be23f24e9aefd59c08dc37bc20b46db",
				"v1.16.13": "bb4d0f045600b883745016416c14533f823d582f4f20df691b7f79a6545b6480",
				"v1.17.0":  "cba12bfe0ee447b06f00813d7d4ba3fbdbf5116eccc4d3291987044f2d6f93c2",
				"v1.17.7":  "00c71ceffa9b50af081d2838b102be49ca224a8aa928f5c948b804af84c58818",
				"v1.17.8":  "4dfd36dbd637b8dca9a7c4e789fb3fe4ca420062c90d3a872ae751dfb9777cb6",
				"v1.17.9":  "4d818e97073113eb1e62bf97d63876757be0f273c47807c09f34511155e25afd",
				"v1.18.5":  "28c1edb2d76f80e70e10fa8cd2a30b9fccc5f003d8b3e853535d8317db7f424a",
				"v1.18.6":  "7b3d6cc019747a7ee5f6cc2b187423daaac4e153140cb290e60d316c3f456430",
				"v1.18.8":  "9046c4086528427462544e1a6dcbe709de4d7ae44d1a155375de330fecd067b1",
				"v1.19.0":  "d4adf1b6b97252025cb2f7febf55daa3f42dc305822e3da133f77fd33071ec2f",
			},
		},
		etcd: {
			amd64: {
				"v3.4.13": "2ac029e47bab752dacdb7b30032f230f49e2f457cbc32e8f555c2210bb5ff107",
			},
			arm64: {
				"v3.4.13": "1934ebb9f9f6501f706111b78e5e321a7ff8d7792d3d96a76e2d01874e42a300",
			},
		},
		helm: {
			amd64: {
				"v3.2.1": "98c57f2b86493dd36ebaab98990e6d5117510f5efbf21c3344c3bdc91a4f947c",
			},
			arm64: {
				"v3.2.1": "20bb9d66e74f618cd104ca07e4525a8f2f760dd6d5611f7d59b6ac574624d672",
			},
		},
		kubecni: {
			amd64: {
				"v0.8.2": "21283754ffb953329388b5a3c52cef7d656d535292bda2d86fcdda604b482f85",
				"v0.8.6": "994fbfcdbb2eedcfa87e48d8edb9bb365f4e2747a7e47658482556c12fd9b2f5",
			},
			arm64: {
				"v0.8.6": "43fbf750c5eccb10accffeeb092693c32b236fb25d919cf058c91a677822c999",
			},
		},
	}
)

func getBinaryCheckSum(name string, arch string, version string) (string, error) {
	avs, ok := fileSha256[name]
	if !ok {
		return "", errors.Wrapf(errors.ErrNotFound, "binary %s", name)
	}
	vs, ok := avs[arch]
	if !ok {
		return "", errors.Wrapf(errors.ErrNotFound, "arch %s", arch)
	}
	s, ok := vs[version]
	if !ok {
		return "", errors.Wrapf(errors.ErrNotFound, "version %s", version)
	}

	return s, nil
}

func GetKubeadmCheckSum(arch, version string) (string, error) {
	return getBinaryCheckSum(kubeadm, arch, version)
}

func GetKubeletCheckSum(arch, version string) (string, error) {
	return getBinaryCheckSum(kubelet, arch, version)
}

func GetKubectlCheckSum(arch, version string) (string, error) {
	return getBinaryCheckSum(kubectl, arch, version)
}

func GetCNICheckSum(arch, version string) (string, error) {
	return getBinaryCheckSum(kubecni, arch, version)
}
