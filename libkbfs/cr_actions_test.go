package libkbfs

import (
	"reflect"
	"testing"
)

func TestCRActionsCollapseNoChange(t *testing.T) {
	al := crActionList{
		&copyUnmergedEntryAction{"old1", "new1", ""},
		&copyUnmergedEntryAction{"old2", "new2", ""},
		&renameUnmergedAction{"old3", "new3"},
		&renameMergedAction{"old4", "new4"},
		&copyUnmergedAttrAction{"old5", "new5", mtimeAttr},
	}

	newList := al.collapse()
	if !reflect.DeepEqual(al, newList) {
		t.Errorf("Collapse returned different list: %v vs %v", al, newList)
	}
}

func TestCRActionsCollapseEntry(t *testing.T) {
	al := crActionList{
		&copyUnmergedAttrAction{"old", "new", mtimeAttr},
		&copyUnmergedEntryAction{"old", "new", ""},
		&renameUnmergedAction{"old", "new"},
	}

	expected := crActionList{
		al[2],
	}

	newList := al.collapse()
	if !reflect.DeepEqual(expected, newList) {
		t.Errorf("Collapse returned unexpected list: %v vs %v",
			expected, newList)
	}

	// change the order
	al = crActionList{al[1], al[2], al[0]}

	newList = al.collapse()
	if !reflect.DeepEqual(expected, newList) {
		t.Errorf("Collapse returned unexpected list: %v vs %v",
			expected, newList)
	}

	// Omit the top action this time
	al = crActionList{al[0], al[2]}
	expected = crActionList{al[0]}

	newList = al.collapse()
	if !reflect.DeepEqual(expected, newList) {
		t.Errorf("Collapse returned unexpected list: %v vs %v",
			expected, newList)
	}
}
