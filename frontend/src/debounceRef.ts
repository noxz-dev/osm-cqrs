import { customRef, Ref, unref } from 'vue';

export type MaybeRef<T> = Ref<T> | T;

/**
 * Delayed ref which triggers the value change after the given delay
 *
 * @param value Value of the ref
 * @param delay The delay when the change should be triggered
 * @param callOutside If it should be called outside
 * @returns The delayed ref with the value
 */
export function debounceRef<T>(value: T, delay: MaybeRef<number> = 200, callOutside: MaybeRef<boolean> = true): Ref<T> {
  let timeout: number;
  return customRef<T>((track, trigger) => {
    return {
      get() {
        track();
        return value;
      },
      set(newValue) {
        clearTimeout(timeout);
        if (unref(callOutside)) {
          value = newValue;
        }
        timeout = setTimeout(() => {
          if (!unref(callOutside)) {
            value = newValue;
          }
          trigger();
        }, unref(delay));
      }
    };
  });
}
