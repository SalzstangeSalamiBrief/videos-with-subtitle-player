import { queryOptions } from '@tanstack/react-query';
import { getFileTreeQuery } from './getFileTreeQuery';
import { getFileTreeSelect } from './getFileTreeSelect';

export const getFileTreeQueryOptions = queryOptions({
  queryKey: ['fileTree'],
  queryFn: getFileTreeQuery,
  select: getFileTreeSelect,
});
