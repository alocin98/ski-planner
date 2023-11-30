const distance = (distance: number) => {
    if (distance < 1000) {
        return `${distance}m`;
    }

    return `${(distance / 1000).toFixed(1)}km`;
}

const time = (timeInSeconds: number) => {
    const hours = Math.floor(timeInSeconds / 3600);
    const minutes = Math.floor((timeInSeconds - hours * 3600) / 60);
    const seconds = timeInSeconds - hours * 3600 - minutes * 60;

    const hoursStr = hours > 0 ? `${hours}h` : '';
    const minutesStr = minutes > 0 ? `${minutes}m` : '';
    const secondsStr = seconds > 0 ? `${seconds}s` : '';

    return `${hoursStr} ${minutesStr} ${secondsStr}`;
}

const elevation = (elevation: number) => {
    return `${elevation}m`;
}

const stringIsoDate = (date: string) => {
    const dateObj = new Date(date);
    const month = dateObj.getMonth() + 1;
    const day = dateObj.getDate();
    const year = dateObj.getFullYear();

    return `${day}/${month}/${year}`;
}

export const Formaters = {
    distance,
    time,
    elevation,
    stringIsoDate
}
