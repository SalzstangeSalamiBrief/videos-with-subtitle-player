import { ApiError } from '@videos-with-subtitle-player/core';

interface IErrorMessageProps {
  error: ApiError | Error;
  message: string;
  description?: string;
}

export function ErrorMessage({
  error,
  message,
  description,
}: IErrorMessageProps) {
  console.error(error);

  if (ApiError.isApiError(error)) {
    return (
      <section className="h-fit rounded-md bg-red-50 p-4 text-red-900">
        <h1 className="text-lg font-bold">
          {error.status}: {error.title}
        </h1>
        <details>
          <summary className="mt-2 cursor-pointer underline">Details</summary>
          <div>
            <p className="whitespace-pre-wrap">{error.detail}</p>
            <p>{error.type}</p>
          </div>
        </details>
      </section>
    );
  }

  // TODO IMPROVE
  return (
    <section className="h-fit rounded-md bg-red-50 p-4 text-red-900">
      <h1 className="text-lg font-bold">{message}</h1>
      {description && <p>{description}</p>}
    </section>
  );
}
