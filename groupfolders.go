package gonextcloud

import "strconv"

type groupFolderBadFormatIDAndGroups struct {
	ID         string            `json:"id"`
	MountPoint string            `json:"mount_point"`
	Groups     map[string]string `json:"groups"`
	Quota      string            `json:"quota"`
	Size       int               `json:"size"`
}

type groupFolderBadFormatGroups struct {
	ID         int               `json:"id"`
	MountPoint string            `json:"mount_point"`
	Groups     map[string]string `json:"groups"`
	Quota      string            `json:"quota"`
	Size       int               `json:"size"`
}

type GroupFolder struct {
	ID         int                        `json:"id"`
	MountPoint string                     `json:"mount_point"`
	Groups     map[string]SharePermission `json:"groups"`
	Quota      int                        `json:"quota"`
	Size       int                        `json:"size"`
}

func (gf *groupFolderBadFormatGroups) FormatGroupFolder() GroupFolder {
	g := GroupFolder{}
	g.ID = gf.ID
	g.MountPoint = gf.MountPoint
	g.Groups = map[string]SharePermission{}
	for k, v := range gf.Groups {
		p, _ := strconv.Atoi(v)
		g.Groups[k] = SharePermission(p)
	}
	q, _ := strconv.Atoi(gf.Quota)
	g.Quota = q
	g.Size = gf.Size
	return g
}

func (gf *groupFolderBadFormatIDAndGroups) FormatGroupFolder() GroupFolder {
	g := GroupFolder{}
	g.ID, _ = strconv.Atoi(gf.ID)
	g.MountPoint = gf.MountPoint
	g.Groups = map[string]SharePermission{}
	for k, v := range gf.Groups {
		p, _ := strconv.Atoi(v)
		g.Groups[k] = SharePermission(p)
	}
	q, _ := strconv.Atoi(gf.Quota)
	g.Quota = q
	g.Size = gf.Size
	return g
}
