package items

import "Avalon/avalon/types"

func ShelfList() error {
	return nil
}

func ShelfOnSell(item *types.ItemShelf) error {
	_, err := engine.Insert(item)
	return err
}

func ShelfUpdate(item *types.ItemShelf) error {
	_, err := engine.Update(item)
	return err
}

func ShelfRemove(id int) error {
	item := &types.ItemShelf{}
	_, err := engine.Id(id).Delete(item)
	return err
}
