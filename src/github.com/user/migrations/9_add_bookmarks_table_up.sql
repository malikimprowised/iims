create table bookmarks (
  id serial primary key  NOT null,
  sessionId int references sessions(id),
  name text
)
