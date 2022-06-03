package WialonRetranslator

type WialonRetranslatorPackage struct {
	PackageSize uint32 `GoBinTool:"Type:Binary ByteOrder:LittleEndian"`
	UID         string
	Time        uint32
}
