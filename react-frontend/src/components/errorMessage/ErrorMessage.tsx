import { Alert } from "antd";

interface IErrorMessageProps {
  error: any;
  message: string;
  description?: string;
}

export function ErrorMessage({
  error,
  message,
  description,
}: IErrorMessageProps) {
  console.error(error);

  return (
    <Alert message={message} description={description} type="error" showIcon />
  );
}
