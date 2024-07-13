import { Flex, Spin } from 'antd';

interface ILoadingSpinnerProps {
  text: string;
}

export function LoadingSpinner({ text }: ILoadingSpinnerProps) {
  return (
    <Flex align="center" vertical>
      <Spin size="large" />
      <p>{text}</p>
    </Flex>
  );
}
