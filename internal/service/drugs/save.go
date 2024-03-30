package drugs

import "context"

func (n *DrugsSrv) SaveDrug(ctx context.Context, userID int64) error {
	d, err := n.GetFromMemory(userID)
	if err != nil {
		return err
	}

	if err := n.checkFields(d); err != nil {
		return err
	}

	return n.drugsEditor.Save(ctx, d)
}
