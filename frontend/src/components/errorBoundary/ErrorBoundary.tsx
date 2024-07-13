import { Alert } from 'antd';
import React, { ReactNode } from 'react';

interface IProps {
  children?: ReactNode;
}

interface IState {
  hasError: boolean;
}

class ErrorBoundary extends React.Component<IProps, IState> {
  constructor(props: IProps) {
    super(props);
    this.state = { hasError: false };
  }

  static getDerivedStateFromError() {
    return { hasError: true };
  }

  componentDidCatch(error: any, errorInfo: any) {
    console.error('error', error);
    console.error('errorInfo', errorInfo);
  }

  render() {
    if (this.state.hasError) {
      return (
        <Alert
          message="Sonething went wrong!"
          description="Please try again later."
          type="error"
          showIcon
        />
      );
    }
    return this.props.children;
  }
}

export default ErrorBoundary;
