interface IErrorMessageProps {
  error: unknown;
  message: string;
  description?: string;
}

export function ErrorMessage({
  error,
  message,
  description,
}: IErrorMessageProps) {
  console.error(error);
  // TODO IMPROVE
  return (
    <section className="h-fit rounded-md bg-red-50 p-4 text-red-900">
      <h1 className="text-lg font-bold">{message}</h1>
      {description && <p>{description}</p>}
    </section>
  );
}
