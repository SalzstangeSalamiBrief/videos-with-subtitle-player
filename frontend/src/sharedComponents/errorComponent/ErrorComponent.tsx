import { useRouter } from '@tanstack/react-router';
import type { ErrorComponentProps } from '@tanstack/react-router';
import { ApiError } from '$models/ApiError';

export function ErrorComponent({ error, reset }: ErrorComponentProps) {
  const router = useRouter();

  const resetButton = (
    <button
      className="btn btn-outline hover:bg-slate-800"
      onClick={() => {
        // Invalidate the route to reload the loader, and reset any router error boundaries
        router.invalidate();
        reset();
      }}
    >
      Retry
    </button>
  );

  if (ApiError.isApiError(error)) {
    return (
      <div role="alert" className="alert alert-error alert-outline">
        <details>
          <summary className="p-4 hover:cursor-pointer">
            {error.status}: {error.title}
          </summary>
          {error.detail}
        </details>
        {resetButton}
      </div>
    );
  }

  return (
    <section
      role="alert"
      className="alert alert-error alert-outline flex flex-col items-start"
    >
      <h2 className="text-lg font-bold">{error.message}</h2>
      {resetButton}
    </section>
  );
}
