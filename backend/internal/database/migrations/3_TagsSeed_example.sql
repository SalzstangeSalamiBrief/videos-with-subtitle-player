DO $$
DECLARE
   tagsToCreate TEXT Array := '{"a", "b"}';
   currentTag TEXT;
BEGIN
    FOREACH currentTag IN ARRAY tagsToCreate
    LOOP
            INSERT INTO tags (name, created_at)
            VALUES (currentTag, CURRENT_TIMESTAMP)
            ON CONFLICT ON CONSTRAINT uni_tags_name DO NOTHING;
    END LOOP;
END
$$;