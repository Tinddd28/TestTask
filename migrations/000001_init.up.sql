
CREATE TABLE IF NOT EXISTS "songs" (
   "id" BIGSERIAL PRIMARY KEY,
   "author_name" VARCHAR NOT NULL,
   "song_name" VARCHAR NOT NULL,
   "release_date" DATE NOT NULL,
   "link" VARCHAR NOT NULL
);

CREATE TABLE IF NOT EXISTS "verses" (
    "id" BIGSERIAL PRIMARY KEY,
    "song_id" BIGINT NOT NULL,
    "text" TEXT NOT NULL,
    CONSTRAINT "fk_song_id" FOREIGN KEY ("song_id") REFERENCES "songs" ("id") ON DELETE CASCADE
);

CREATE INDEX "idx_song_id" ON "verses" ("song_id");
CREATE INDEX "idx_song" ON "songs" ("song");
CREATE INDEX "idx_group" ON "songs" ("group");
CREATE INDEX "idx_release_date" ON "songs" ("release_date");
