// Package item provides models and functions for working with items
package item

import (
	"github.com/AevtJJ/idmangler/block"
	"github.com/AevtJJ/idmangler/types"
)

// ShinyProperty represents a single shiny property on an item
type ShinyProperty struct {
	ID    byte
	Value int64
}

// Item represents a game item with its properties
type Item struct {
	// Basic properties
	Name     string
	ItemType types.ItemType

	// Optional properties
	Element        *types.Element
	AttackSpeed    *types.AttackSpeed
	PowderSlots    byte
	Powders        []types.Powder
	GearType       *types.CraftedGearType
	Durability     *int
	ConsumableType *types.ConsumableType
	Effects        []types.Effect

	// Shiny properties
	ShinyProps []ShinyProperty
}

// ToBlocks converts the item to a series of blocks for encoding
func (i *Item) ToBlocks() []block.AnyBlock {
	blocks := make([]block.AnyBlock, 0)

	// Add start block
	blocks = append(blocks, block.NewStartData(types.Version1))

	// Add type block
	blocks = append(blocks, block.NewTypeData(i.ItemType))

	// Add name block
	blocks = append(blocks, block.NewNameData(i.Name))

	// Add powder data if we have powders or powder slots
	if i.PowderSlots > 0 || len(i.Powders) > 0 {
		blocks = append(blocks, block.NewPowderData(i.PowderSlots, i.Powders))
	}

	// Add shiny properties if any
	for _, shiny := range i.ShinyProps {
		blocks = append(blocks, block.NewShinyData(shiny.ID, shiny.Value))
	}

	// Always end with the end block
	blocks = append(blocks, block.NewEndData())

	return blocks
}

// FromBlocks creates an Item from a series of blocks
func FromBlocks(blocks []block.AnyBlock) (*Item, error) {
	item := &Item{}

	// Process each block
	for _, b := range blocks {
		switch block := b.(type) {
		case *block.TypeData:
			item.ItemType = block.ItemType

		case *block.NameData:
			item.Name = block.Name

		case *block.PowderData:
			item.PowderSlots = block.PowderSlots
			item.Powders = block.Powders

		case *block.ShinyData:
			if item.ShinyProps == nil {
				item.ShinyProps = make([]ShinyProperty, 0)
			}
			item.ShinyProps = append(item.ShinyProps, ShinyProperty{
				ID:    block.ID,
				Value: block.Value,
			})
		}
	}

	return item, nil
}

// Encode encodes the item into an ID string
func (i *Item) Encode() (string, error) {
	blocks := i.ToBlocks()

	encoder := block.NewItemEncoder()
	return encoder.EncodeBlocks(blocks)
}

// NewBasicItem creates a new item with the specified name and type
func NewBasicItem(name string, itemType types.ItemType) *Item {
	return &Item{
		Name:     name,
		ItemType: itemType,
	}
}

// AddPowder adds a powder to the item
func (i *Item) AddPowder(powder types.Powder) {
	if i.Powders == nil {
		i.Powders = make([]types.Powder, 0)
	}
	i.Powders = append(i.Powders, powder)
}

// SetPowderSlots sets the number of powder slots on the item
func (i *Item) SetPowderSlots(slots byte) {
	i.PowderSlots = slots
}

// AddShinyProperty adds a shiny property to the item
func (i *Item) AddShinyProperty(id byte, value int64) {
	if i.ShinyProps == nil {
		i.ShinyProps = make([]ShinyProperty, 0)
	}
	i.ShinyProps = append(i.ShinyProps, ShinyProperty{
		ID:    id,
		Value: value,
	})
}
