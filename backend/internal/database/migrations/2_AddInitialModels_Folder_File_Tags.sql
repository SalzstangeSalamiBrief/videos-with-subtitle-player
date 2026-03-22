CREATE TABLE IF NOT EXISTS folder_nodes (
    id                       BIGSERIAL PRIMARY KEY,
    created_at               TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at               TIMESTAMP NOT NULL DEFAULT NOW(),
    folder_id                UUID NOT NULL,
    name                     VARCHAR NOT NULL,
    path                     VARCHAR NOT NULL,
    CONSTRAINT uni_folder_nodes_folder_id UNIQUE (folder_id),
    CONSTRAINT uni_folder_nodes_name      UNIQUE (name),
    CONSTRAINT uni_folder_nodes_path      UNIQUE (path),
    thumbnail_id             VARCHAR NOT NULL DEFAULT '',
    low_quality_thumbnail_id VARCHAR NOT NULL DEFAULT '',
    parent_folder_id         UUID NULL REFERENCES folder_nodes (folder_id) ON UPDATE CASCADE ON DELETE CASCADE
);

CREATE INDEX IF NOT EXISTS idx_folder_nodes_parent_folder_id ON folder_nodes (parent_folder_id);

CREATE TABLE IF NOT EXISTS "fileNodes" (
    id                       BIGSERIAL PRIMARY KEY,
    created_at               TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at               TIMESTAMP NOT NULL DEFAULT NOW(),
    file_id                  UUID,
    path                     VARCHAR NOT NULL,
    name                     VARCHAR,
    CONSTRAINT "uni_fileNodes_path" UNIQUE (path),
    type                     file_type NOT NULL,
    associated_audio_file_id UUID,
    low_quality_image_id     UUID,
    parent_folder_id         UUID REFERENCES folder_nodes (folder_id) ON UPDATE CASCADE ON DELETE CASCADE
);

CREATE INDEX IF NOT EXISTS idx_file_nodes_parent_folder_id ON "fileNodes" (parent_folder_id);

CREATE TABLE IF NOT EXISTS tags (
    id         BIGSERIAL PRIMARY KEY,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    name       VARCHAR NOT NULL,
    CONSTRAINT uni_tags_name UNIQUE (name)
);

CREATE TABLE IF NOT EXISTS folder_node_to_tags (
    folder_node_id        BIGINT NOT NULL REFERENCES folder_nodes (id) ON UPDATE CASCADE ON DELETE CASCADE,
    tag_id                BIGINT NOT NULL REFERENCES tags (id) ON UPDATE CASCADE ON DELETE CASCADE,
    PRIMARY KEY (folder_node_id, tag_id)
);
