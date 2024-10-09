drop index "idx_song_id";
drop index "idx_song";
drop index "idx_group";
drop index "idx_release_date";
alter table "verses"
    drop constraint "fk_song_id";

drop table if exists "verses";
drop table if exists "songs";