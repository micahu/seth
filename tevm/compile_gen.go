package tevm

// NOTE: THIS FILE WAS PRODUCED BY THE
// MSGP CODE GENERATION TOOL (github.com/tinylib/msgp)
// DO NOT EDIT

import (
	"github.com/tinylib/msgp/msgp"
)

// DecodeMsg implements msgp.Decodable
func (z *CompiledBundle) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "filenames":
			var zb0002 uint32
			zb0002, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Filenames) >= int(zb0002) {
				z.Filenames = (z.Filenames)[:zb0002]
			} else {
				z.Filenames = make([]string, zb0002)
			}
			for za0001 := range z.Filenames {
				z.Filenames[za0001], err = dc.ReadString()
				if err != nil {
					return
				}
			}
		case "sources":
			var zb0003 uint32
			zb0003, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Sources) >= int(zb0003) {
				z.Sources = (z.Sources)[:zb0003]
			} else {
				z.Sources = make([]string, zb0003)
			}
			for za0002 := range z.Sources {
				z.Sources[za0002], err = dc.ReadString()
				if err != nil {
					return
				}
			}
		case "contracts":
			var zb0004 uint32
			zb0004, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Contracts) >= int(zb0004) {
				z.Contracts = (z.Contracts)[:zb0004]
			} else {
				z.Contracts = make([]CompiledContract, zb0004)
			}
			for za0003 := range z.Contracts {
				var zb0005 uint32
				zb0005, err = dc.ReadMapHeader()
				if err != nil {
					return
				}
				for zb0005 > 0 {
					zb0005--
					field, err = dc.ReadMapKeyPtr()
					if err != nil {
						return
					}
					switch msgp.UnsafeString(field) {
					case "name":
						z.Contracts[za0003].Name, err = dc.ReadString()
						if err != nil {
							return
						}
					case "code":
						z.Contracts[za0003].Code, err = dc.ReadBytes(z.Contracts[za0003].Code)
						if err != nil {
							return
						}
					case "sourcemap":
						z.Contracts[za0003].Sourcemap, err = dc.ReadString()
						if err != nil {
							return
						}
					default:
						err = dc.Skip()
						if err != nil {
							return
						}
					}
				}
			}
		case "warnings":
			var zb0006 uint32
			zb0006, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Warnings) >= int(zb0006) {
				z.Warnings = (z.Warnings)[:zb0006]
			} else {
				z.Warnings = make([]string, zb0006)
			}
			for za0004 := range z.Warnings {
				z.Warnings[za0004], err = dc.ReadString()
				if err != nil {
					return
				}
			}
		default:
			err = dc.Skip()
			if err != nil {
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *CompiledBundle) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 4
	// write "filenames"
	err = en.Append(0x84, 0xa9, 0x66, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x73)
	if err != nil {
		return err
	}
	err = en.WriteArrayHeader(uint32(len(z.Filenames)))
	if err != nil {
		return
	}
	for za0001 := range z.Filenames {
		err = en.WriteString(z.Filenames[za0001])
		if err != nil {
			return
		}
	}
	// write "sources"
	err = en.Append(0xa7, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73)
	if err != nil {
		return err
	}
	err = en.WriteArrayHeader(uint32(len(z.Sources)))
	if err != nil {
		return
	}
	for za0002 := range z.Sources {
		err = en.WriteString(z.Sources[za0002])
		if err != nil {
			return
		}
	}
	// write "contracts"
	err = en.Append(0xa9, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x73)
	if err != nil {
		return err
	}
	err = en.WriteArrayHeader(uint32(len(z.Contracts)))
	if err != nil {
		return
	}
	for za0003 := range z.Contracts {
		// map header, size 3
		// write "name"
		err = en.Append(0x83, 0xa4, 0x6e, 0x61, 0x6d, 0x65)
		if err != nil {
			return err
		}
		err = en.WriteString(z.Contracts[za0003].Name)
		if err != nil {
			return
		}
		// write "code"
		err = en.Append(0xa4, 0x63, 0x6f, 0x64, 0x65)
		if err != nil {
			return err
		}
		err = en.WriteBytes(z.Contracts[za0003].Code)
		if err != nil {
			return
		}
		// write "sourcemap"
		err = en.Append(0xa9, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x6d, 0x61, 0x70)
		if err != nil {
			return err
		}
		err = en.WriteString(z.Contracts[za0003].Sourcemap)
		if err != nil {
			return
		}
	}
	// write "warnings"
	err = en.Append(0xa8, 0x77, 0x61, 0x72, 0x6e, 0x69, 0x6e, 0x67, 0x73)
	if err != nil {
		return err
	}
	err = en.WriteArrayHeader(uint32(len(z.Warnings)))
	if err != nil {
		return
	}
	for za0004 := range z.Warnings {
		err = en.WriteString(z.Warnings[za0004])
		if err != nil {
			return
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *CompiledBundle) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 4
	// string "filenames"
	o = append(o, 0x84, 0xa9, 0x66, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x73)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Filenames)))
	for za0001 := range z.Filenames {
		o = msgp.AppendString(o, z.Filenames[za0001])
	}
	// string "sources"
	o = append(o, 0xa7, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Sources)))
	for za0002 := range z.Sources {
		o = msgp.AppendString(o, z.Sources[za0002])
	}
	// string "contracts"
	o = append(o, 0xa9, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x73)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Contracts)))
	for za0003 := range z.Contracts {
		// map header, size 3
		// string "name"
		o = append(o, 0x83, 0xa4, 0x6e, 0x61, 0x6d, 0x65)
		o = msgp.AppendString(o, z.Contracts[za0003].Name)
		// string "code"
		o = append(o, 0xa4, 0x63, 0x6f, 0x64, 0x65)
		o = msgp.AppendBytes(o, z.Contracts[za0003].Code)
		// string "sourcemap"
		o = append(o, 0xa9, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x6d, 0x61, 0x70)
		o = msgp.AppendString(o, z.Contracts[za0003].Sourcemap)
	}
	// string "warnings"
	o = append(o, 0xa8, 0x77, 0x61, 0x72, 0x6e, 0x69, 0x6e, 0x67, 0x73)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Warnings)))
	for za0004 := range z.Warnings {
		o = msgp.AppendString(o, z.Warnings[za0004])
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *CompiledBundle) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "filenames":
			var zb0002 uint32
			zb0002, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Filenames) >= int(zb0002) {
				z.Filenames = (z.Filenames)[:zb0002]
			} else {
				z.Filenames = make([]string, zb0002)
			}
			for za0001 := range z.Filenames {
				z.Filenames[za0001], bts, err = msgp.ReadStringBytes(bts)
				if err != nil {
					return
				}
			}
		case "sources":
			var zb0003 uint32
			zb0003, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Sources) >= int(zb0003) {
				z.Sources = (z.Sources)[:zb0003]
			} else {
				z.Sources = make([]string, zb0003)
			}
			for za0002 := range z.Sources {
				z.Sources[za0002], bts, err = msgp.ReadStringBytes(bts)
				if err != nil {
					return
				}
			}
		case "contracts":
			var zb0004 uint32
			zb0004, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Contracts) >= int(zb0004) {
				z.Contracts = (z.Contracts)[:zb0004]
			} else {
				z.Contracts = make([]CompiledContract, zb0004)
			}
			for za0003 := range z.Contracts {
				var zb0005 uint32
				zb0005, bts, err = msgp.ReadMapHeaderBytes(bts)
				if err != nil {
					return
				}
				for zb0005 > 0 {
					zb0005--
					field, bts, err = msgp.ReadMapKeyZC(bts)
					if err != nil {
						return
					}
					switch msgp.UnsafeString(field) {
					case "name":
						z.Contracts[za0003].Name, bts, err = msgp.ReadStringBytes(bts)
						if err != nil {
							return
						}
					case "code":
						z.Contracts[za0003].Code, bts, err = msgp.ReadBytesBytes(bts, z.Contracts[za0003].Code)
						if err != nil {
							return
						}
					case "sourcemap":
						z.Contracts[za0003].Sourcemap, bts, err = msgp.ReadStringBytes(bts)
						if err != nil {
							return
						}
					default:
						bts, err = msgp.Skip(bts)
						if err != nil {
							return
						}
					}
				}
			}
		case "warnings":
			var zb0006 uint32
			zb0006, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Warnings) >= int(zb0006) {
				z.Warnings = (z.Warnings)[:zb0006]
			} else {
				z.Warnings = make([]string, zb0006)
			}
			for za0004 := range z.Warnings {
				z.Warnings[za0004], bts, err = msgp.ReadStringBytes(bts)
				if err != nil {
					return
				}
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *CompiledBundle) Msgsize() (s int) {
	s = 1 + 10 + msgp.ArrayHeaderSize
	for za0001 := range z.Filenames {
		s += msgp.StringPrefixSize + len(z.Filenames[za0001])
	}
	s += 8 + msgp.ArrayHeaderSize
	for za0002 := range z.Sources {
		s += msgp.StringPrefixSize + len(z.Sources[za0002])
	}
	s += 10 + msgp.ArrayHeaderSize
	for za0003 := range z.Contracts {
		s += 1 + 5 + msgp.StringPrefixSize + len(z.Contracts[za0003].Name) + 5 + msgp.BytesPrefixSize + len(z.Contracts[za0003].Code) + 10 + msgp.StringPrefixSize + len(z.Contracts[za0003].Sourcemap)
	}
	s += 9 + msgp.ArrayHeaderSize
	for za0004 := range z.Warnings {
		s += msgp.StringPrefixSize + len(z.Warnings[za0004])
	}
	return
}

// DecodeMsg implements msgp.Decodable
func (z *CompiledContract) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "name":
			z.Name, err = dc.ReadString()
			if err != nil {
				return
			}
		case "code":
			z.Code, err = dc.ReadBytes(z.Code)
			if err != nil {
				return
			}
		case "sourcemap":
			z.Sourcemap, err = dc.ReadString()
			if err != nil {
				return
			}
		default:
			err = dc.Skip()
			if err != nil {
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *CompiledContract) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 3
	// write "name"
	err = en.Append(0x83, 0xa4, 0x6e, 0x61, 0x6d, 0x65)
	if err != nil {
		return err
	}
	err = en.WriteString(z.Name)
	if err != nil {
		return
	}
	// write "code"
	err = en.Append(0xa4, 0x63, 0x6f, 0x64, 0x65)
	if err != nil {
		return err
	}
	err = en.WriteBytes(z.Code)
	if err != nil {
		return
	}
	// write "sourcemap"
	err = en.Append(0xa9, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x6d, 0x61, 0x70)
	if err != nil {
		return err
	}
	err = en.WriteString(z.Sourcemap)
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *CompiledContract) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 3
	// string "name"
	o = append(o, 0x83, 0xa4, 0x6e, 0x61, 0x6d, 0x65)
	o = msgp.AppendString(o, z.Name)
	// string "code"
	o = append(o, 0xa4, 0x63, 0x6f, 0x64, 0x65)
	o = msgp.AppendBytes(o, z.Code)
	// string "sourcemap"
	o = append(o, 0xa9, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x6d, 0x61, 0x70)
	o = msgp.AppendString(o, z.Sourcemap)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *CompiledContract) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "name":
			z.Name, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "code":
			z.Code, bts, err = msgp.ReadBytesBytes(bts, z.Code)
			if err != nil {
				return
			}
		case "sourcemap":
			z.Sourcemap, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *CompiledContract) Msgsize() (s int) {
	s = 1 + 5 + msgp.StringPrefixSize + len(z.Name) + 5 + msgp.BytesPrefixSize + len(z.Code) + 10 + msgp.StringPrefixSize + len(z.Sourcemap)
	return
}

// DecodeMsg implements msgp.Decodable
func (z *Source) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "Filename":
			z.Filename, err = dc.ReadString()
			if err != nil {
				return
			}
		case "Body":
			z.Body, err = dc.ReadString()
			if err != nil {
				return
			}
		default:
			err = dc.Skip()
			if err != nil {
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z Source) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 2
	// write "Filename"
	err = en.Append(0x82, 0xa8, 0x46, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65)
	if err != nil {
		return err
	}
	err = en.WriteString(z.Filename)
	if err != nil {
		return
	}
	// write "Body"
	err = en.Append(0xa4, 0x42, 0x6f, 0x64, 0x79)
	if err != nil {
		return err
	}
	err = en.WriteString(z.Body)
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z Source) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 2
	// string "Filename"
	o = append(o, 0x82, 0xa8, 0x46, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65)
	o = msgp.AppendString(o, z.Filename)
	// string "Body"
	o = append(o, 0xa4, 0x42, 0x6f, 0x64, 0x79)
	o = msgp.AppendString(o, z.Body)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Source) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "Filename":
			z.Filename, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "Body":
			z.Body, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z Source) Msgsize() (s int) {
	s = 1 + 9 + msgp.StringPrefixSize + len(z.Filename) + 5 + msgp.StringPrefixSize + len(z.Body)
	return
}
