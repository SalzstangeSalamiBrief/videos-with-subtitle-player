DO $$
DECLARE
    level_record RECORD;
    item_record RECORD;
    num_tags INT;
    selected_tags BIGINT[];
    tag_id BIGINT;
BEGIN
    -- Iterate over each unique directory level (parent path)
    FOR level_record IN
        SELECT DISTINCT regexp_replace(path, '/[^/]*$', '') AS parent_path
        FROM "file_tree_items"
    LOOP
        -- Pick a random number of tags from 0 to 12 for this level
        num_tags := floor(random() * 13)::INT;

        -- Select num_tags random tag ids
        SELECT ARRAY(
            SELECT id
            FROM tags
            ORDER BY random()
            LIMIT num_tags
        ) INTO selected_tags;

        -- Assign these tags to all fileTreeItems on this level
        FOR item_record IN
            SELECT id
            FROM "file_tree_items"
            WHERE regexp_replace(path, '/[^/]*$', '') = level_record.parent_path
        LOOP
            FOREACH tag_id IN ARRAY selected_tags
            LOOP
                INSERT INTO "file_tree_item_to_tags" (file_tree_item_id, tag_id)
                VALUES (item_record.id, tag_id)
                ON CONFLICT DO NOTHING;
            END LOOP;
        END LOOP;

    END LOOP;
END $$;
