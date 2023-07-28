package serializer

import (
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"
	"pcbook/pb"
	"pcbook/sample"
	"testing"
)

func TestFileSerializer(t *testing.T) {
	t.Parallel()

	binPath := "../tmp/laptop.bin"
	jsonPath := "../tmp/laptop.json"

	laptop1 := sample.NewLaptop()
	err := WriteProtobufToBinaryFile(laptop1, binPath)
	require.NoError(t, err)

	laptop2 := &pb.Laptop{}
	err = ReadProtobufFromBinaryFile(binPath, laptop2)
	require.True(t, proto.Equal(laptop1, laptop2))

	err = WriteProtobufToJSONFile(laptop1, jsonPath)
	require.NoError(t, err)
}
