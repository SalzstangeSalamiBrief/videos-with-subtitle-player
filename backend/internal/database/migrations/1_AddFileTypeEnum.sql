DO $$
DECLARE
    fileTypeEnumDoesNotExist BOOLEAN := (SELECT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'file_type') IS NOT true);
BEGIN
    IF fileTypeEnumDoesNotExist THEN 
		CREATE TYPE file_type AS ENUM ('Video', 'Audio', 'Subtitle', 'Image', 'Unknown');
	END IF;
END
$$;