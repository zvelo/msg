package main

import (
	"github.com/gogo/protobuf/vanity"
	"github.com/gogo/protobuf/vanity/command"
)

func main() {
	req := command.Read()
	files := req.GetProtoFile()
	files = vanity.FilterFiles(files, vanity.NotGoogleProtobufDescriptorProto)

	vanity.ForEachFile(files, vanity.TurnOnMarshalerAll)   // gofast
	vanity.ForEachFile(files, vanity.TurnOnSizerAll)       // gofast
	vanity.ForEachFile(files, vanity.TurnOnUnmarshalerAll) // gofast

	vanity.ForEachFieldInFilesExcludingExtensions(vanity.OnlyProto2(files), vanity.TurnOffNullableForNativeTypesWithoutDefaultsOnly) // gogofaster
	vanity.ForEachFile(files, vanity.TurnOffGoUnrecognizedAll)                                                                       // gogofaster

	vanity.ForEachFile(files, vanity.TurnOffGoEnumPrefixAll)   // gogoslick
	vanity.ForEachFile(files, vanity.TurnOffGoEnumStringerAll) // gogoslick
	vanity.ForEachFile(files, vanity.TurnOnEnumStringerAll)    // gogoslick

	vanity.ForEachFile(files, vanity.TurnOnEqualAll)       // gogoslick
	vanity.ForEachFile(files, vanity.TurnOnGoStringAll)    // gogoslick
	vanity.ForEachFile(files, vanity.TurnOffGoStringerAll) // gogoslick
	vanity.ForEachFile(files, vanity.TurnOnStringerAll)    // gogoslick

	vanity.ForEachFile(files, vanity.TurnOnVerboseEqualAll) // zvelo
	vanity.ForEachFile(files, vanity.TurnOnCompareAll)      // zvelo

	resp := command.Generate(req)
	command.Write(resp)
}
