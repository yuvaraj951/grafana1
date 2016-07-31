package models

import (
  "time"
  "errors"
)

type Process struct {
  OrgId    int64
  ProcessId       int64
  ProcessName     string
  ParentProcessName string
  Created         time.Time
  Updated         time.Time
  UpdatedBy       string
}
// ---------------------
// COMMANDS
type AddProcessCommand struct {
  OrgId    int64            `json:"orgId"`
  ProcessId       int64   `json:"processId"`
  ProcessName     string    `json:"processName" binding:"Required"`
  ParentProcessName string     `json:"parentProcessName" binding:"Required"`
  UpdatedBy       string      `json:"updatedBy" binding:"Required"`

}
// Typed errors
var (

  ErrProcessAlreadyAdded = errors.New("User is already added to organization")
ErrOrgFound=errors.New("User is already added to oation")
)
// ----------------------
// QUERIES

type GetProcessQuery struct {
  OrgId    int64
  Result []*ProcessDTO
}
// ----------------------
// Projections and DTOs
type ProcessDTO struct {
  OrgId    int64            `json:"orgId"`
  ProcessId       int64   `json:"processId"`
  ProcessName     string    `json:"processName"`
  ParentProcessName string     `json:"parentProcessName"`
  Created         time.Time   `json:"created"`
  Updated         time.Time   `json:"updated"`
  UpdatedBy       string      `json:"updatedBy"`


}

type GetProcessByProcessIdQuery  struct {
  OrgId    int64
  Result *Process

}
