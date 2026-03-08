import { defineConfig } from 'orval';

export default defineConfig({
  petstore: {
    output: {
      mode: 'tags-split',
      target: 'src/client/generated/client.ts',
      schemas: 'src/client/generated/',
      client: 'fetch',
      baseUrl: 'http://localhost:3000',
      mock: false,
      override: {
        enumGenerationType: 'enum',
      },
    },
    input: {
      target: 'http://localhost:3000/openapi.yaml',
    },
  },
});
