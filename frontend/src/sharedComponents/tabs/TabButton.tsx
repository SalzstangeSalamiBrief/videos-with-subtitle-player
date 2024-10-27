import { getTabId, getTabPanelId } from './Tabs';

interface ITabButtonProps {
  label: string;
  isActive: boolean;
  index: number;
  onClick: () => void;
}

export function TabButton({
  label,
  isActive,
  onClick,
  index,
}: ITabButtonProps) {
  return (
    <button
      key={index}
      role="tab"
      id={getTabId(index)}
      aria-selected={isActive ? 'true' : 'false'}
      aria-controls={getTabPanelId(index)}
      onClick={onClick}
      className={`px-4 py-2 ${isActive ? 'bg-fuchsia-800 hover:bg-fuchsia-700' : 'bg-slate-800 hover:bg-slate-700'} rounded-md`}
    >
      {label}
    </button>
  );
}
