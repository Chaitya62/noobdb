package page

import "errors"

const PAGE_SIZE = 4096
const INVALID_PAGE_ID = 4294967295

type Page interface {
	GetData() []byte
	GetPageId() uint32
	SetData([]byte) error
	SetPageId(page_id uint32)
}

type PageImpl struct {
	_data      [PAGE_SIZE]byte
	_page_id   uint32
	_dirty     bool
	_pin_count uint64
}

func (p *PageImpl) GetData() []byte {
	return p._data[:]
}

func (p *PageImpl) GetPageId() uint32 {
	return p._page_id
}

func (p *PageImpl) SetData(d []byte) error {
	if len(d) != PAGE_SIZE {
		return errors.New("Invalid Data slice")
	}
	copy(p._data[:], d)
	p._page_id = (uint32(d[0]) | (uint32(d[1]) << 8) | (uint32(d[2]) << 16) | (uint32(d[3]) << 24))
	p._dirty = true
	return nil
}

func (p *PageImpl) ResetMemory() error {
	copy(p._data[:], make([]byte, PAGE_SIZE))
	p._page_id = 0
	return nil
}

func (p *PageImpl) SetPageId(page_id uint32) {
	p._page_id = page_id
	p._data[0] = byte(page_id)
	p._data[1] = byte(page_id >> 8)
	p._data[2] = byte(page_id >> 16)
	p._data[3] = byte(page_id >> 24)
	p._dirty = true
}

func InvalidPage() Page {
	// the first four bits are no equal to page_id
	return &PageImpl{[PAGE_SIZE]byte{}, INVALID_PAGE_ID, false, 0}
}
