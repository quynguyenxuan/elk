package elk

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGroups(t *testing.T) {
	gs := groups{}

	gs.Add("group")
	require.Len(t, gs, 1)
	require.True(t, gs.HasGroup("group"))
	require.False(t, gs.HasGroup("none"))

	gs.Add("group_1", "group_2")
	require.Len(t, gs, 3)
	require.True(t, gs.HasGroup("group"))
	require.True(t, gs.HasGroup("group_1"))
	require.True(t, gs.HasGroup("group_2"))
	require.False(t, gs.HasGroup("none"))

	require.False(t, gs.Match(groups{"none", "nobody"}))
	require.True(t, gs.Match(groups{"group", "nobody"}))

	require.Equal(t, `groups:"group_one,GROUP_two,group:3"`, groups{"group_one", "GROUP_two", "group:3"}.StructTag())

	require.True(t, groups{}.Equal(groups{}))
	require.True(t, groups{"a"}.Equal(groups{"a"}))
	require.True(t, groups{"a", "b"}.Equal(groups{"a", "b"}))
	require.False(t, groups{"a"}.Equal(groups{}))
	require.False(t, groups{"a"}.Equal(groups{"b"}))

	require.Equal(t, groups{}.Hash(), groups{}.Hash())
	require.Equal(t, groups{"a"}.Hash(), groups{"a"}.Hash())
	require.Equal(t, groups{"a", "b"}.Hash(), groups{"a", "b"}.Hash())
	require.NotEqual(t, groups{"a", "b"}.Hash(), groups{"ab"}.Hash())
	require.NotEqual(t, groups{"a"}.Hash(), groups{"b"}.Hash())
}
