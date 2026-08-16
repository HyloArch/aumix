import type { PropsWithChildren } from "react";

type PageProps = {
  hidden: boolean;
};

export default function Page({
  hidden,
  children,
}: PropsWithChildren<PageProps>) {
  return (
    <div className="page" hidden={hidden}>
      {children}
    </div>
  );
}
