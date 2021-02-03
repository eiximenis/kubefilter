package main

import (
	"flag"
	"strings"
)

type commandLineArgs struct {
	logLevel int
	removeNil bool
	removeEmpty  bool
	removeOwnerRefs bool
	additionalKeys []string
}

func (clargs* commandLineArgs) HasDebugLevel() bool {
	return clargs.logLevel == 8
}

func parseCommandLines() *commandLineArgs {
	removeNil := flag.Bool("remove-null", true, "Remove null values")
	removeEmpty := flag.Bool("remove-empty", true, "Remove empty objects")
	removeOwner := flag.Bool("remove-owner-refs", false, "Remove metadata.ownerReferences")
	logLevel := flag.Int("log-level", 0, "Log level (0 none, 8 debug)")
	additionalKeys := flag.String("remove-keys", "", "Additional keys to remove (use full name like metadata.name), comma-separated")
	flag.Parse()
	clargs := commandLineArgs{
		logLevel:    		*logLevel,
		removeNil:   		*removeNil,
		removeEmpty: 		*removeEmpty,
		removeOwnerRefs:	*removeOwner,
		additionalKeys: 	strings.Split(*additionalKeys, ","),
	}

	return &clargs
}

