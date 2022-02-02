package iodb

import (
	"io"

	"github.com/TechPay-io/sirius-base/common/bigendian"
	"github.com/TechPay-io/sirius-base/kvdb"

	"github.com/TechPay-io/go-photon/utils/ioread"
)

func Write(writer io.Writer, db kvdb.Iteratee) error {
	it := db.NewIterator(nil, nil)
	defer it.Release()
	for it.Next() {
		_, err := writer.Write(bigendian.Uint32ToBytes(uint32(len(it.Key()))))
		if err != nil {
			return err
		}
		_, err = writer.Write(it.Key())
		if err != nil {
			return err
		}
		_, err = writer.Write(bigendian.Uint32ToBytes(uint32(len(it.Value()))))
		if err != nil {
			return err
		}
		_, err = writer.Write(it.Value())
		if err != nil {
			return err
		}
	}
	return nil
}

func Read(reader io.Reader, batch kvdb.Batch) error {
	defer batch.Reset()
	var lenB [4]byte
	for {
		err := ioread.ReadAll(reader, lenB[:])
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		lenKey := bigendian.BytesToUint32(lenB[:])
		key := make([]byte, lenKey)
		err = ioread.ReadAll(reader, key)
		if err != nil {
			return err
		}

		err = ioread.ReadAll(reader, lenB[:])
		if err != nil {
			return err
		}

		lenValue := bigendian.BytesToUint32(lenB[:])
		value := make([]byte, lenValue)
		err = ioread.ReadAll(reader, value)
		if err != nil {
			return err
		}

		err = batch.Put(key, value)
		if err != nil {
			return err
		}
		if batch.ValueSize() > kvdb.IdealBatchSize {
			err = batch.Write()
			if err != nil {
				return err
			}
			batch.Reset()
		}
	}
	return batch.Write()
}
