import type { MouseEvent, PropsWithChildren } from "react";
import "./Popup.css";
import { clsx } from "../../util";

type PopupProps = {
  isOpen: boolean;
  onClose: () => void;
  className?: string;
};

export default function Popup({
  isOpen,
  onClose,
  className,
  children,
}: PropsWithChildren<PopupProps>) {
  return (
    <>
      {isOpen && (
        <div className={clsx("popup", className)} onClick={onClose}>
          <div
            className="popup-container"
            onClick={(e: MouseEvent) => e.stopPropagation()}
          >
            {children}
          </div>
        </div>
      )}
    </>
  );
}
