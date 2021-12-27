package fs

var (
	// docker run --name node --rm -it node:12-alpine sh
	// yarn init -y
	// yarn add promise jquery
	// yarn list | grep -E -o "\S+@[^\^~]\S+" | awk -F@ 'NR>0 {printf("{\""$1"\", \""$2"\", \"\"},\n")}'
	yarnNormal = YarnLockfile{
		"asap@~2.0.6": &LockfileEntry{
			Version:              "2.0.6",
			Integrity:            "sha1-5QNHYR1+aQlDIIu9r+vLwvuGbUY=",
			Resolved:             "https://registry.yarnpkg.com/asap/-/asap-2.0.6.tgz#e50347611d7e690943208bbdafebcbc2fb866d46",
			Dependencies:         nil,
			OptionalDependencies: nil,
		},
		"jquery@^3.4.1": &LockfileEntry{
			Version:              "3.4.1",
			Integrity:            "sha512-36+AdBzCL+y6qjw5Tx7HgzeGCzC81MDDgaUP8ld2zhx58HdqXGoBd+tHdrBMiyjGQs0Hxs/MLZTu/eHNJJuWPw==",
			Resolved:             "https://registry.yarnpkg.com/jquery/-/jquery-3.4.1.tgz#714f1f8d9dde4bdfa55764ba37ef214630d80ef2",
			Dependencies:         nil,
			OptionalDependencies: nil,
		},
		"promise@^8.0.3": &LockfileEntry{
			Version:   "8.0.3",
			Integrity: "sha512-HeRDUL1RJiLhyA0/grn+PTShlBAcLuh/1BJGtrvjwbvRDCTLLMEz9rOGCV+R3vHY4MixIuoMEd9Yq/XvsTPcjw==",
			Resolved:  "https://registry.yarnpkg.com/promise/-/promise-8.0.3.tgz#f592e099c6cddc000d538ee7283bb190452b0bf6",
			Dependencies: map[string]string{
				"asap": "~2.0.6",
			},
			OptionalDependencies: nil,
		},
	}

	yarnNormalV2 = YarnLockfile{
		"__metadata": &LockfileEntry{
			Version:  "4",
			Resolved: "",
		},
		"asap@npm:~2.0.6": &LockfileEntry{
			Version:              "2.0.6",
			Integrity:            "",
			Resolved:             "",
			Dependencies:         nil,
			OptionalDependencies: nil,
		},
	}

	yarnNormalWindows = YarnLockfile{
		"asap@~2.0.6": &LockfileEntry{
			Version:              "2.0.6",
			Integrity:            "sha1-5QNHYR1+aQlDIIu9r+vLwvuGbUY=",
			Resolved:             "https://registry.yarnpkg.com/asap/-/asap-2.0.6.tgz#e50347611d7e690943208bbdafebcbc2fb866d46",
			Dependencies:         nil,
			OptionalDependencies: nil,
		},
		"jquery@^3.4.1": &LockfileEntry{
			Version:              "3.4.1",
			Integrity:            "sha512-36+AdBzCL+y6qjw5Tx7HgzeGCzC81MDDgaUP8ld2zhx58HdqXGoBd+tHdrBMiyjGQs0Hxs/MLZTu/eHNJJuWPw==",
			Resolved:             "https://registry.yarnpkg.com/jquery/-/jquery-3.4.1.tgz#714f1f8d9dde4bdfa55764ba37ef214630d80ef2",
			Dependencies:         nil,
			OptionalDependencies: nil,
		},
		"promise@^8.0.3": &LockfileEntry{
			Version:   "8.0.3",
			Integrity: "sha512-HeRDUL1RJiLhyA0/grn+PTShlBAcLuh/1BJGtrvjwbvRDCTLLMEz9rOGCV+R3vHY4MixIuoMEd9Yq/XvsTPcjw==",
			Resolved:  "https://registry.yarnpkg.com/promise/-/promise-8.0.3.tgz#f592e099c6cddc000d538ee7283bb190452b0bf6",
			Dependencies: map[string]string{
				"asap": "~2.0.6",
			},
			OptionalDependencies: nil,
		},
	}
)
