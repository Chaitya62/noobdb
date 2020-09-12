package page

const PAGE_SIZE = 4096

type Page interface {
	GetData() [PAGE_SIZE]byte
	GetPageId() uint32
	ResetMemory() error
}

type PageImpl struct {
	_data    [PAGE_SIZE]byte
	_page_id uint32
}

func (p PageImpl) GetData() [PAGE_SIZE]byte {
	return p._data
}

func (p PageImpl) GetPageId() uint32 {
	return p._page_id
}

func (p *PageImpl) ResetMemory() error {
	copy(p._data[:], make([]byte, 4096))
	return nil
}
