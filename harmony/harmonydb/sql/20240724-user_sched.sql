CREATE TABLE harmony_task_user (
  task_id INTEGER PRIMARY KEY, 
  owner TEXT NOT NULL, 
  expiration INTEGER NOT NULL
);