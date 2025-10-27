import { useQueryErrorResetBoundary } from '@tanstack/react-query';
import { useRouter, type ErrorComponentProps } from '@tanstack/react-router';
import { ApiError } from '@videos-with-subtitle-player/core';
import { useEffect } from 'react';

export function ErrorComponent({ error, reset }: ErrorComponentProps) {
  console.log('ERROR COMPONENT RENDERED WITH ERROR:', error);
  const router = useRouter();
  const queryErrorResetBoundary = useQueryErrorResetBoundary();

  useEffect(() => {
    // Reset the query error boundary
    // queryErrorResetBoundary.reset();
    // reset();
  }, [queryErrorResetBoundary, reset]);

  const resetButton = (
    <button
      onClick={() => {
        // Invalidate the route to reload the loader, and reset any router error boundaries
        router.invalidate();
        queryErrorResetBoundary.reset();
        reset();
      }}
    >
      retry
    </button>
  );

  if (ApiError.isApiError(error)) {
    return (
      <div className="mx-auto w-full max-w-120 rounded bg-red-800 px-4">
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
    <div>
      {error.message}
      {resetButton}
    </div>
  );
}
