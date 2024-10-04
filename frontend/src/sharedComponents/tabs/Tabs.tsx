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
  const navigate = useNavigate({ from: Route.fullPath });
  const labelId = useId();
  const searchParams: RootSearchParams = useSearch({ strict: false });
  const activeTabIndex = getActiveTabIndex(searchParams.activeTab, tabs.length);
  if (activeTabIndex !== searchParams.activeTab) {
    navigate({
      search: () => ({
        activeTab: activeTabIndex,
      }),
    });
  }

  if (!tabs.length) {
    return null;
  }

  const activeTab = tabs[activeTabIndex];
  return (
    <section aria-labelledby={labelId} className="h-full pb-8">
      <h1 id={labelId} className="sr-only">
        {label}
      </h1>

      <div role="tablist" className="flex flex-col gap-8 h-full">
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
          className="flex-grow"
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

function getActiveTabIndex(
  input: number | undefined,
  numberOfTabs: number,
): number {
  if (input === undefined) {
    return 0;
  }

  if (Number.isNaN(input)) {
    return 0;
  }

  if (input < 0) {
    return 0;
  }

  if (numberOfTabs < input) {
    return numberOfTabs - 1;
  }

  return input;
}

function getTabId(index: number): string {
  return `tab-${index}`;
}

function getTabPanelId(index: number): string {
  return `tabpanel-${index}`;
}
