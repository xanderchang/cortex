package db

import (
	"github.com/Sirupsen/logrus"
	"github.com/weaveworks/cortex/configs"
)

// traced adds logrus trace lines on each db call
type traced struct {
	d DB
}

func (t traced) trace(name string, args ...interface{}) {
	logrus.Debugf("%s: %#v", name, args)
}

func (t traced) GetOrgConfig(orgID configs.OrgID, subsystem configs.Subsystem) (cfg configs.ConfigView, err error) {
	defer func() { t.trace("GetOrgConfig", orgID, subsystem, cfg, err) }()
	return t.d.GetOrgConfig(orgID, subsystem)
}

func (t traced) SetOrgConfig(orgID configs.OrgID, subsystem configs.Subsystem, cfg configs.Config) (err error) {
	defer func() { t.trace("SetOrgConfig", orgID, subsystem, cfg, err) }()
	return t.d.SetOrgConfig(orgID, subsystem, cfg)
}

func (t traced) GetAllOrgConfigs(subsystem configs.Subsystem) (cfgs map[configs.OrgID]configs.ConfigView, err error) {
	defer func() { t.trace("GetAllOrgConfigs", subsystem, cfgs, err) }()
	return t.d.GetAllOrgConfigs(subsystem)
}

func (t traced) GetOrgConfigs(subsystem configs.Subsystem, since configs.ID) (cfgs map[configs.OrgID]configs.ConfigView, err error) {
	defer func() { t.trace("GetOrgConfigs", subsystem, since, cfgs, err) }()
	return t.d.GetOrgConfigs(subsystem, since)
}

func (t traced) Close() (err error) {
	defer func() { t.trace("Close", err) }()
	return t.d.Close()
}
