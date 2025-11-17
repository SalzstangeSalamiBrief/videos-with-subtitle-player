import { getTabId, getTabPanelId } from './tabUtilities';

interface ITabButtonProps {
  label: string;
  isActive: boolean;
  index: number;
  onClick: (selectedTabIndex: number) => void;
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
      onClick={() => onClick(index)}
      // className={`tab px-4 py-2 ${isActive ? 'bg-fuchsia-800 hover:bg-fuchsia-700' : 'bg-slate-800 hover:bg-slate-700'} rounded-md`}
      className={`tab ${isActive ? 'tab-active' : ''}`}
    >
      {label}
    </button>
  );
}
