import styles from './LoadingSpinner.module.css';

interface ILoadingSpinnerProps {
  text: string;
}

export function LoadingSpinner({ text }: ILoadingSpinnerProps) {
  return (
    <div className="flex flex-col items-center">
      <div className={styles.loader} />
      <p>{text}</p>
    </div>
  );
}
