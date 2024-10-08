export type Image = {
    name: string;
    namespace: string;
    description: string;
    pullCount: number;
};

export type SearchedImage = {
    repo_name: string,
    short_description: string,
    star_count: number,
    pull_count: number,
    repo_owner: string,
    is_automated: boolean,
    is_official: boolean
}

export type AppParams = {
    appName: string;
    imageName: string;
}
export const fetchDockerImages = async (searchTerm: string): Promise<any> => {
    const formData = new URLSearchParams();
    formData.append("searchTerm", searchTerm);

    const req = new Request("?/searchDockerHubRepos", {
        method: "POST",
        headers: {
            "Content-Type": "application/x-www-form-urlencoded",
        },
        body: formData.toString(),
    });

    const res = await fetch(req);
    if (!res.ok) throw new Error(`Error: ${res.statusText}`);

    const data = await res.json();
    const parsedData = JSON.parse(data.data);
    return JSON.parse(parsedData[0]).results;
};

export interface MetricData {
    cpu_stats: {
        cpu_usage: {
            total_usage: number;
            percpu_usage: number[];
            usage_in_kernelmode: number;
            usage_in_usermode: number;
        };
        system_cpu_usage: number;
        online_cpus: number;
    };
    memory_stats: {
        usage: number;
        max_usage: number;
        stats: {
            rss: number;
            cache: number;
        };
    };
}

export function getUsedMemory(MetricData: MetricData): number {
    const RAMUsed =
        MetricData.memory_stats.usage - MetricData.memory_stats.stats.cache;
    return RAMUsed / 1048576
}

let preCPUSystemUsage: number;
export function getPercentageOfCPUUsed(metricData: MetricData): number {
    let preCPUUsage: number = 0
    metricData.cpu_stats.cpu_usage.percpu_usage.forEach((consumptionThread: number) => { preCPUUsage += consumptionThread })

    const deltaSystemCPUUsage = metricData.cpu_stats.system_cpu_usage - preCPUSystemUsage
    if (preCPUSystemUsage != undefined) {
        const PercentageOfCPUUsed = (preCPUUsage / deltaSystemCPUUsage) * metricData.cpu_stats.cpu_usage.percpu_usage.length * 100
        preCPUSystemUsage = metricData.cpu_stats.system_cpu_usage
        return PercentageOfCPUUsed
    }
    preCPUSystemUsage = metricData.cpu_stats.system_cpu_usage
    return 0
}