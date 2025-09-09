// Package idmangler provides functionality for encoding and decoding Wynntils ID strings
package idmangler

import (
	"github.com/AevtJJ/idmangler/block"
	"github.com/AevtJJ/idmangler/item"
	"github.com/AevtJJ/idmangler/types"
)

// Version represents the current version of the idmangler library
const Version = "0.1.0"

// EncodeItem encodes an item represented as a series of blocks into an ID string
func EncodeItem(blocks []block.AnyBlock) (string, error) {
	encoder := block.NewItemEncoder()
	return encoder.EncodeBlocks(blocks)
}

// DecodeItem decodes an ID string into a series of blocks
func DecodeItem(idString string) ([]block.AnyBlock, error) {
	decoder := block.NewItemDecoder()
	return decoder.DecodeString(idString)
}

// EncodeItemObject encodes an Item object into an ID string
func EncodeItemObject(item *item.Item) (string, error) {
	blocks := item.ToBlocks()
	return EncodeItem(blocks)
}

// DecodeItemObject decodes an ID string into an Item object
func DecodeItemObject(idString string) (*item.Item, error) {
	blocks, err := DecodeItem(idString)
	if err != nil {
		return nil, err
	}

	return item.FromBlocks(blocks)
}

// CreateBasicItem creates a basic item with the minimum required blocks
func CreateBasicItem(name string, itemType types.ItemType) []block.AnyBlock {
	startBlock := block.NewStartData(types.Version1)
	typeBlock := block.NewTypeData(itemType)
	nameBlock := block.NewNameData(name)
	endBlock := block.NewEndData()

	return []block.AnyBlock{startBlock, typeBlock, nameBlock, endBlock}
}
