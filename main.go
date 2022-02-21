package main

import (
	registry "refsiverdur.org/node-pensioner-manager/v2/registry"
)

func main() {
	age := registry.GetPackageAge("@vleesbrood/unbg")
	println(age)
}
