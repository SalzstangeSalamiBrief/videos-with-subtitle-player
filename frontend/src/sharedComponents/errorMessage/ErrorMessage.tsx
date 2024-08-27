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
    <section className="bg-red-50 text-red-900 p-4 rounded-md">
      <h1 className="font-bold text-lg">{message}</h1>
      <p>{description}</p>
    </section>
  );
}
