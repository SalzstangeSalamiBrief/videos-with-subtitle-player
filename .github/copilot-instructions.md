# Copilot Instructions for Videos with Subtitle Player

## Architecture

- **Backend**: Go with modular architecture
  - API endpoints in `backend/internal/routes/`
  - Services in `backend/pkg/services/`
  - Models in `backend/pkg/models/`
  - Utilities in `backend/pkg/utilities/`
- **Frontend**: Multi-framework approach with workspace structure
  - Nuxt.js app in `frontend/apps/nuxt/`
  - React app in `frontend/apps/react/`
  - Svelte app in `frontend/apps/svelte/`
  - Shared packages in `frontend/packages/`

## Coding Standards & Conventions

### Go Backend

- Follow Go standard conventions (gofmt, golint)
- Keep handlers thin, business logic in services
- Use proper error handling with custom error types
- Structure: `cmd/` for main applications, `internal/` for private code, `pkg/` for public libraries

#### Testing

- Test files should be co-located with source files
- For tests use the model in _backend\pkg\models\testModels.go_ to ensure consistency and create multiple tests without overhead

### Frontend

- Use TypeScript for all new JavaScript/TypeScript code
- Follow framework-specific conventions (Nuxt, React, Svelte)
- Shared code goes in `packages/core/`
- Use ESLint and Prettier configurations from `packages/eslintBase/` and `packages/prettierBase/`
- Use pnpm as package manager

### File Organization

- Keep related files close together
- Do not barrel exports (`index.ts`) for clean imports
- Group by feature, not by file type
- Test files should be co-located with source files or in dedicated test directories

## Key Features & Context

- Video playback with subtitle synchronization
- Multi-format support for video and subtitle files
- WebP image processing capabilities
- Audio file handling (MP3 support evident from test content)

## Development Guidelines

1. **API Development**:

   - RESTful endpoints following Go best practices
   - Proper HTTP status codes and error responses
   - Input validation and sanitization

2. **Frontend Development**:

   - Component-based architecture
   - Responsive design principles
   - Accessibility considerations for media players

3. **Testing**:

   - Unit tests for services and utilities
   - Integration tests for API endpoints
   - Component testing for frontend features

## Dependencies & Tools

- Go modules for backend dependency management
- pnpm workspaces for frontend monorepo
- Docker for containerization

## Security Considerations

- Validate file uploads and types
- Sanitize file paths to prevent directory traversal
- Implement proper CORS for frontend-backend communication
- Consider rate limiting for media endpoints

When suggesting code, prioritize:

1. Type safety (TypeScript/Go type system)
2. Error handling and edge cases
3. Performance optimizations for media handling
4. Clean, maintainable code structure
5. Proper resource management and cleanup
