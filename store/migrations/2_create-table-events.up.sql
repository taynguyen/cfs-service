CREATE TABLE events (
  id VARCHAR(56),
  number VARCHAR(56),
  agency_id VARCHAR(56),
  type_code VARCHAR(24),
  created_time DATETIME(3),
  dispatch_time DATETIME(3),
  response_code VARCHAR(56),
  inserted_at DATETIME(3),

  PRIMARY KEY (id)
);

CREATE INDEX events_number ON events (number);
CREATE INDEX events_agency ON events (agency_id);
CREATE INDEX events_createdtime ON events (agency_id, created_time);
CREATE INDEX events_agency_dispathtime ON events (agency_id, dispatch_time);