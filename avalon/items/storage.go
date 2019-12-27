package items

import "Avalon/avalon/types"

func GetItemById(id int) (*types.ItemInfo, error) {
	item := &types.ItemInfo{}
	_, err := engine.Id(id).Get(item)
	return item, err
}

func AddItem(item *types.ItemInfo) error {
	_, err := engine.Insert(item)
	return err
}

func UpdateItem(item *types.ItemInfo) error {
	_, err := engine.Update(item)
	return err
}

func DeleteItemById(id int) error {
	item := &types.ItemInfo{}
	_, err := engine.Id(id).Delete(item)
	return err
}
