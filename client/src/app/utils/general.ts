export const getRoundedDownHour = (): number => {
    const now = new Date();

    const roundedDownHour = new Date(
        now.getFullYear(),
        now.getMonth(),
        now.getDate(),
        now.getHours()
    ).getTime(); // Start of the current hour as timestamp

    return roundedDownHour;
}

export const trimToFirstDelimiter = (input: string): string => {
  return input.split(/[,\s;]/)[0];
}