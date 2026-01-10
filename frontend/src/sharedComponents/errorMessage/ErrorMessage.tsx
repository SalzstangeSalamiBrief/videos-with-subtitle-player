import { ApiError } from '$models/ApiError';

interface IErrorMessageProps {
  error: ApiError | Error | string;
  description?: string;
}

export function ErrorMessage(props: IErrorMessageProps) {
  console.error(props.error);

  const { content, title } = getContent(props);

  return (
    <section
      role="alert"
      className="alert alert-error alert-outline flex flex-col items-start"
    >
      <h2 className="text-lg font-bold">{title}</h2>
      {content && content}
    </section>
  );
}

function getContent({ error, description }: IErrorMessageProps) {
  if (ApiError.isApiError(error)) {
    return {
      title: `${error.status}: ${error.title}`,
      content: (
        <details>
          <summary className="mt-2 cursor-pointer underline">Details</summary>
          <div>
            <p className="whitespace-pre-wrap">{error.detail}</p>
            <p>{error.type}</p>
          </div>
        </details>
      ),
    };
  }

  if (typeof error === 'string') {
    return {
      title: error,
      content: null,
    };
  }

  return {
    title: error?.toString(),
    content: description ? <p>{description}</p> : null,
  };
}
