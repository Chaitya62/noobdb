package page

const PAGE_SIZE = 4096
const INVALID_PAGE_ID = 4294967295

type Page interface {
	GetData() [PAGE_SIZE]byte
	GetPageId() uint32
	SetData([]byte) error
	SetPageId(page_id uint32)
}

type PageImpl struct {
	_data    [PAGE_SIZE]byte
	_page_id uint32
}

func (p *PageImpl) GetData() [PAGE_SIZE]byte {
	return p._data
}

func (p *PageImpl) GetPageId() uint32 {
	return p._page_id
}

func (p *PageImpl) SetData(d []byte) error {
	copy(p._data[:], d)
	return nil
}

func (p *PageImpl) ResetMemory() error {
	copy(p._data[:], make([]byte, 4096))
	return nil
}

func (p *PageImpl) SetPageId(page_id uint32) {
	p._page_id = page_id
}

func InvalidPage() Page {
	return &PageImpl{[PAGE_SIZE]byte{}, INVALID_PAGE_ID}
}
