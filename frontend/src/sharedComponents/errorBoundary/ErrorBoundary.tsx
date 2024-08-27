import React, { ErrorInfo, ReactNode } from 'react';

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

  componentDidCatch(error: unknown, errorInfo: ErrorInfo) {
    console.error('error', error);
    console.error('errorInfo', errorInfo);
  }

  render() {
    if (this.state.hasError) {
      return (
        <section className="bg-red-800">
          <h1>Something went wrong!</h1>
          <p>Please try again later.</p>
        </section>
      );
    }
    return this.props.children;
  }
}

export default ErrorBoundary;
