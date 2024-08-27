import { useNavigate, useSearch } from '@tanstack/react-router';
import { useId } from 'react';
import { Route } from '../../routes';
import { RootSearchParams } from '../../routes/folders/_folderLayout';

export interface ITab {
  label: string;
  content: JSX.Element;
}

interface ITabsProps {
  tabs: ITab[];
  label: string;
}
// TODO MOVE TO FEATURE
export function Tabs({ tabs, label }: ITabsProps) {
  const labelId = useId();
  const searchParams: RootSearchParams = useSearch({ strict: false });
  const activeTabIndex = searchParams.activeTab ?? 0;
  const navigate = useNavigate({ from: Route.fullPath });

  if (!tabs.length) {
    return null;
  }

  const activeTab = tabs[activeTabIndex];
  return (
    <section aria-labelledby={labelId}>
      <h1 id={labelId} className="sr-only">
        {label}
      </h1>

      <div role="tablist" className="grid gap-8">
        <div role="presentation" className="flex gap-2">
          {tabs.map((tab, index) => {
            const isActive = index === activeTabIndex;

            return (
              <button
                key={index}
                role="tab"
                id={getTabId(index)}
                aria-selected={isActive ? 'true' : 'false'}
                aria-controls={getTabPanelId(index)}
                onClick={() =>
                  navigate({
                    search: () => ({ activeTab: index }),
                  })
                }
                className={`px-4 py-2 ${isActive ? 'bg-fuchsia-800 hover:bg-fuchsia-700' : 'bg-slate-800 hover:bg-slate-700'} rounded-md`}
              >
                {tab.label}
              </button>
            );
          })}
        </div>

        <div
          id={getTabPanelId(activeTabIndex)}
          role="tabpanel"
          tabIndex={0}
          aria-labelledby={getTabId(activeTabIndex)}
        >
          {activeTab.content}
        </div>
      </div>
    </section>
  );
}

function getTabId(index: number): string {
  return `tab-${index}`;
}

function getTabPanelId(index: number): string {
  return `tabpanel-${index}`;
}
