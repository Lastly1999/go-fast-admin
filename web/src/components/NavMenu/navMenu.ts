
export interface MenuItem {
    id: number;
    label: string;
    value: string;
    path: string;
    pId: number;
    pName: string;
}

export interface MenuInfo extends MenuItem {
    children: MenuItem[]
}