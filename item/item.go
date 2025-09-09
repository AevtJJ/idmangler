package item

import (
	"github.com/AevtJJ/idmangler/block"
	"github.com/AevtJJ/idmangler/encoding"
	"github.com/AevtJJ/idmangler/types"
)

// Item represents a generic Wynn item that can be encoded into and decoded from ID strings
type Item struct {
	// Name of the item
	Name string
	// ItemType is the type of the item
	ItemType types.ItemType
	// PowderSlots is the number of powder slots the item has
	PowderSlots int
	// Powders is a list of powders applied to the item
	Powders []types.Powder
	// Identifications is a list of stats on the item
	Identifications []*types.Stat
	// ShinyProps contains shiny properties of the item
	ShinyProps []ShinyProp
	// Rerolls is the number of times the item has been rerolled
	Rerolls byte
	// Extended encoding mode for identifications
	ExtendedEncoding bool
}

// ShinyProp represents a shiny property on an item
type ShinyProp struct {
	ID    byte
	Value int64
}

// NewBasicItem creates a new basic item with the given name and type
func NewBasicItem(name string, itemType types.ItemType) *Item {
	return &Item{
		Name:             name,
		ItemType:         itemType,
		Powders:          make([]types.Powder, 0),
		Identifications:  make([]*types.Stat, 0),
		ShinyProps:       make([]ShinyProp, 0),
		ExtendedEncoding: true,
	}
}

// SetPowderSlots sets the number of powder slots on the item
func (i *Item) SetPowderSlots(slots int) {
	i.PowderSlots = slots
}

// AddPowder adds a powder to the item with the given element and tier
func (i *Item) AddPowder(element int, tier int) (types.Powder, error) {
	if len(i.Powders) >= i.PowderSlots {
		return types.Powder{}, &encoding.EncodeError{
			Type: encoding.ErrTooManyPowders,
		}
	}

	elem, err := types.ElementFromID(byte(element))
	if err != nil {
		return types.Powder{}, &encoding.EncodeError{
			Type:    encoding.ErrBadElement,
			Details: element,
		}
	}

	powder, err := types.NewPowder(elem, byte(tier))
	if err != nil {
		return types.Powder{}, &encoding.EncodeError{
			Type:    encoding.ErrBadPowderTier,
			Details: tier,
		}
	}

	i.Powders = append(i.Powders, powder)
	return powder, nil
}

// AddShinyProperty adds a shiny property to the item
func (i *Item) AddShinyProperty(id byte, value int64) {
	i.ShinyProps = append(i.ShinyProps, ShinyProp{ID: id, Value: value})
}

// AddIdentification adds an identification stat to the item
func (i *Item) AddIdentification(kind byte, baseValue int32, roll byte) {
	basePtr := &baseValue
	i.Identifications = append(i.Identifications, types.NewStat(kind, basePtr, types.NewValueRoll(roll)))
}

// AddPreIdentifiedStat adds a pre-identified stat to the item
func (i *Item) AddPreIdentifiedStat(kind byte, baseValue int32) {
	basePtr := &baseValue
	i.Identifications = append(i.Identifications, types.NewStat(kind, basePtr, types.NewPreIdentifiedRoll()))
}

// ToBlocks converts the item to a slice of blocks
func (i *Item) ToBlocks() []block.AnyBlock {
	blocks := make([]block.AnyBlock, 0)

	// Always start with StartData
	startData := &block.StartData{Version: types.Version1}
	blocks = append(blocks, startData)

	// Add TypeData
	typeData := &block.TypeData{ItemType: i.ItemType}
	blocks = append(blocks, typeData)

	// Add NameData if name is present
	if i.Name != "" {
		nameData := &block.NameData{Name: i.Name}
		blocks = append(blocks, nameData)
	}

	// Add IdentificationData if identifications are present
	if len(i.Identifications) > 0 {
		identData := &block.IdentificationData{
			Identifications:  i.Identifications,
			ExtendedEncoding: i.ExtendedEncoding,
		}
		blocks = append(blocks, identData)
	}

	// Add PowderData if powders are present
	if i.PowderSlots > 0 || len(i.Powders) > 0 {
		powderData := &block.PowderData{
			PowderSlots: byte(i.PowderSlots),
			Powders:     i.Powders,
		}
		blocks = append(blocks, powderData)
	}

	// Add RerollData if the item has been rerolled
	if i.Rerolls > 0 {
		rerollData := &block.RerollData{Rerolls: i.Rerolls}
		blocks = append(blocks, rerollData)
	}

	// Add ShinyData if shiny properties are present
	for _, prop := range i.ShinyProps {
		shinyData := &block.ShinyData{
			ID:    prop.ID,
			Value: prop.Value,
		}
		blocks = append(blocks, shinyData)
	}

	// Always end with EndData
	endData := &block.EndData{}
	blocks = append(blocks, endData)

	return blocks
}

// FromBlocks creates an item from a slice of blocks
func FromBlocks(blocks []block.AnyBlock) (*Item, error) {
	if len(blocks) == 0 {
		return nil, &encoding.DecoderError{
			ErrorData: &encoding.DecodeError{Type: encoding.ErrNoStartBlockFound},
		}
	}

	item := &Item{
		Powders:          make([]types.Powder, 0),
		Identifications:  make([]*types.Stat, 0),
		ShinyProps:       make([]ShinyProp, 0),
		ExtendedEncoding: true,
		Rerolls:          0,
	}

	// Process each block
	for _, b := range blocks {
		switch block := b.(type) {
		case *block.StartData:
			// Already handled by the decoder

		case *block.TypeData:
			item.ItemType = block.ItemType

		case *block.NameData:
			item.Name = block.Name

		case *block.IdentificationData:
			item.Identifications = block.Identifications
			item.ExtendedEncoding = block.ExtendedEncoding

		case *block.PowderData:
			item.PowderSlots = int(block.PowderSlots)
			item.Powders = block.Powders

		case *block.RerollData:
			item.Rerolls = block.Rerolls

		case *block.ShinyData:
			item.ShinyProps = append(item.ShinyProps, ShinyProp{
				ID:    block.ID,
				Value: block.Value,
			})

		case *block.EndData:
			// End of data, stop processing
			break
		}
	}

	return item, nil
}
